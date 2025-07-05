package mainsrc

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/s-macke/AdventureAI/src/zmachine"
	"log"
	"os"
	"sync"
	"time"
)

type SessionState struct {
	zm        *zmachine.ZMachine
	StartTime time.Time
	input     string
}

func (session *SessionState) RunUntilInput() string {
	output := ""
	for !session.zm.Done && !session.zm.IsNextZRead() {
		session.zm.InterpretInstruction()
		if session.zm.Output.Len() > 0 {
			if session.zm.WindowId == 0 {
				output += session.zm.Output.String()
			}
			session.zm.Output.Reset()
		}
	}
	return output
}

type SessionManager struct {
	sessions map[string]*SessionState
	filename string
	mutex    sync.RWMutex
}

func NewSessionManager(filename string) *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*SessionState),
		filename: filename,
	}
}

func (sm *SessionManager) CreateSession(sessionID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	zm := Init(sm.filename)
	sm.sessions[sessionID] = &SessionState{
		zm:        zm,
		StartTime: time.Now(),
	}
	zm.Input = func() string {
		session, exists := sm.GetSession(sessionID)
		if !exists {
			return ""
		}
		input := session.input
		session.input = ""
		return input
	}

}

func (sm *SessionManager) GetSession(sessionID string) (*SessionState, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	session, exists := sm.sessions[sessionID]
	return session, exists
}

func (sm *SessionManager) RemoveSession(sessionID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.sessions, sessionID)
}

func createSessionAwareTool(sm *SessionManager) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

		sessionID := server.ClientSessionFromContext(ctx).SessionID()
		session, exists := sm.GetSession(sessionID)
		if !exists {
			return nil, fmt.Errorf("invalid session")
		}
		log.Printf("Received request: %v", request)
		cmd, err := request.RequireString("cmd")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		session.input = cmd
		output := session.RunUntilInput()
		return mcp.NewToolResultText(output), nil
	}
}

func StartServer(filename string) {
	sessionManager := NewSessionManager(filename)
	hooks := &server.Hooks{}
	hooks.AddOnRegisterSession(func(ctx context.Context, session server.ClientSession) {
		sessionManager.CreateSession(session.SessionID())
		log.Printf("Session %s started", session.SessionID())
	})
	hooks.AddOnUnregisterSession(func(ctx context.Context, session server.ClientSession) {
		sessionManager.RemoveSession(session.SessionID())
		log.Printf("Session %s ended", session.SessionID())
	})

	port := os.Getenv("MCP_PORT")
	if port == "" {
		port = "8080"
	}

	s := server.NewMCPServer(
		"AIAdventure",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
		server.WithHooks(hooks),
	)

	tool := mcp.NewTool("execute",
		mcp.WithDescription("executes a command"),
		mcp.WithString("cmd",
			mcp.Required(),
			mcp.Description("simple command, typically one or two words"),
		),
	)

	s.AddTool(tool, createSessionAwareTool(sessionManager))

	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("mcp"),
	)

	fmt.Println("Starting MCP server on port " + port)
	if err := httpServer.Start(":" + port); err != nil {
		log.Fatal(err)
	}

}
