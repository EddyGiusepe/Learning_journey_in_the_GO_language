/**
 * Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro
 *
 * App.jsx
 * =======
 * React chat UI for the Go ADK agent in first_agent.go.
 * Creates a session, sends questions to POST /api/run, and
 * shows the Researcher landing page until the first message.
 *
 * RUN (from the web/ folder, with the Go API already running):
 * ---
 * npm run dev
 */
import { useEffect, useRef, useState } from "react";

const APP_NAME = "researcher";

function extractReply(events) {
  if (!Array.isArray(events)) {
    return typeof events === "string" ? events : JSON.stringify(events);
  }

  const chunks = [];
  for (const event of events) {
    const parts = event?.content?.parts ?? [];
    for (const part of parts) {
      if (typeof part.text === "string" && part.text.trim()) {
        chunks.push(part.text);
      }
    }
  }

  return chunks.join("\n").trim() || "O agente não devolveu texto.";
}

export default function App() {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [busy, setBusy] = useState(false);
  const [ready, setReady] = useState(false);
  const [error, setError] = useState("");
  const ids = useRef({
    userId: "local-user",
    sessionId: crypto.randomUUID(),
  });
  const listRef = useRef(null);

  function sessionUrl() {
    const { userId, sessionId } = ids.current;
    return `/api/apps/${APP_NAME}/users/${userId}/sessions/${sessionId}`;
  }

  async function startSession() {
    const created = await fetch(sessionUrl(), {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({}),
    });
    if (created.ok) {
      return;
    }

    // O ADK devolve 500 quando a sessão já existe (Strict Mode / retry).
    const existing = await fetch(sessionUrl());
    if (existing.ok) {
      return;
    }

    throw new Error(
      `Não consegui criar a sessão (${created.status}). Rode o agente com: go run first_agent.go web api`,
    );
  }

  useEffect(() => {
    let cancelled = false;

    async function connect() {
      for (let attempt = 1; attempt <= 8; attempt += 1) {
        try {
          await startSession();
          if (!cancelled) {
            setReady(true);
            setError("");
          }
          return;
        } catch (err) {
          if (cancelled) {
            return;
          }
          if (attempt === 8) {
            setError(err.message);
            return;
          }
          await new Promise((resolve) => setTimeout(resolve, 700));
        }
      }
    }

    connect();
    return () => {
      cancelled = true;
    };
  }, []);

  useEffect(() => {
    listRef.current?.scrollTo({ top: listRef.current.scrollHeight });
  }, [messages, busy]);

  const examples = [
    "In which country is Machu Picchu and what does that name mean?",
    "What are the latest news about IA and agents?",
  ];

  async function sendMessage(event) {
    event.preventDefault();
    await ask(input);
  }

  async function ask(raw) {
    const text = raw.trim();
    if (!text || busy || !ready) {
      return;
    }

    setInput("");
    setMessages((current) => [...current, { role: "user", text }]);
    setBusy(true);
    setError("");

    try {
      const { userId, sessionId } = ids.current;
      const response = await fetch("/api/run", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          appName: APP_NAME,
          userId,
          sessionId,
          newMessage: {
            role: "user",
            parts: [{ text }],
          },
        }),
      });

      const payload = await response.json();
      if (!response.ok) {
        throw new Error(payload?.detail || payload?.message || `Erro ${response.status}`);
      }

      setMessages((current) => [
        ...current,
        { role: "assistant", text: extractReply(payload) },
      ]);
    } catch (err) {
      setError(err.message);
      setMessages((current) => [
        ...current,
        { role: "assistant", text: `Falha ao falar com o agente: ${err.message}` },
      ]);
    } finally {
      setBusy(false);
    }
  }

  return (
    <div className="app">
      <header>
        <h1>ADK Researcher</h1>
        <p className="powered">
          Powered by ADK + Go
          <span className="pill">
            <i />
            Researcher
          </span>
        </p>
      </header>

      {error ? <p className="banner">{error}</p> : null}

      <div className="chat" ref={listRef}>
        {messages.length === 0 ? (
          <section className="landing">
            <div className="robot" aria-hidden="true">
              🤖
            </div>
            <h2>ADK Researcher</h2>
            <p className="lead">
              I can search topics with the Researcher agent.
            </p>
            <div className="cards">
              <article className="card search">
                <p className="card-title">
                  <i />
                  Researcher
                </p>
                <p>
                  Searches the web for current topics, news and concepts in a detailed way.
                </p>
              </article>
            </div>
            <p className="route">
              Send a question and the agent will search for you.
            </p>
            <div className="examples">
              {examples.map((example) => (
                <button
                  key={example}
                  type="button"
                  className="example"
                  disabled={!ready || busy}
                  onClick={() => ask(example)}
                >
                  Example: {example}
                </button>
              ))}
            </div>
          </section>
        ) : (
          messages.map((message, index) => (
            <article key={index} className={`row ${message.role}`}>
              {message.role === "assistant" ? (
                <p className="meta">
                  ASSISTANT
                  <span className="pill">
                    <i />
                    RESEARCHER
                  </span>
                </p>
              ) : (
                <p className="meta">YOU</p>
              )}
              <p className="bubble">{message.text}</p>
            </article>
          ))
        )}
        {busy ? <p className="typing">The agent is searching…</p> : null}
      </div>

      <form onSubmit={sendMessage}>
        <input
          value={input}
          onChange={(event) => setInput(event.target.value)}
          placeholder={
            ready
              ? "Ask about a topic to search…"
              : "Connecting to the agent…"
          }
          disabled={!ready || busy}
        />
        <button type="submit" disabled={!ready || busy || !input.trim()}>
          Send
        </button>
      </form>

      <footer>
        <p>Developed with ❤️</p>
        <p>
          LinkedIn: Senior Data Scientist/AI Engineering.:{" "}
          <a
            href="https://www.linkedin.com/in/eddy-giusepe-chirinos-isidro-phd-85a43a42/"
            target="_blank"
            rel="noreferrer"
          >
            Dr. Eddy Giusepe Chirinos Isidro
          </a>
        </p>
      </footer>
    </div>
  );
}
