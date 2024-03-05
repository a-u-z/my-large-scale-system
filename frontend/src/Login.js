// Login.js
import React, { useState } from 'react';

function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loginMessage, setLoginMessage] = useState('');

  const handleLogin = async () => {
    try {
      const response = await fetch('http://localhost:8080/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        setLoginMessage('登入成功');
        // 在實際應用中，你可能會在這裡處理登入成功的相應邏輯，例如導向到其他頁面
      } else {
        setLoginMessage('帳號或密碼錯誤');
      }
    } catch (error) {
      setLoginMessage('發生錯誤');
    }
  };

  const handleKeyDown = event => {
    if (event.key === 'Enter') {
      handleLogin();
    }
  };

  return (
    <div>
      <h1>登入系統</h1>
      <div>
        <label>帳號：</label>
        <input
          type="text"
          value={username}
          onChange={e => setUsername(e.target.value)}
        />
      </div>
      <div>
        <label>密碼：</label>
        <input
          type="password"
          value={password}
          onChange={e => setPassword(e.target.value)}
          onKeyDown={handleKeyDown}
        />
      </div>
      <button onClick={handleLogin}>登入</button>
      <p>{loginMessage}</p>
    </div>
  );
}

export default Login;
