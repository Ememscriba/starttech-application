import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [health, setHealth] = useState('checking...');

  useEffect(() => {
    fetch('/health')
      .then(res => res.json())
      .then(data => setHealth(data.status))
      .catch(() => setHealth('offline'));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>StartTech</h1>
        <p>Full Stack Application</p>
        <div className="status">
          <span>API Status: </span>
          <span className={health === 'ok' ? 'online' : 'offline'}>
            {health}
          </span>
        </div>
      </header>
    </div>
  );
}

export default App;
