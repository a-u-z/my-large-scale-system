// Register.js
import React, { useState } from 'react';

function Register() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [registerMessage, setRegisterMessage] = useState('');

  const handleRegister = async () => {
    try {
      // 處理註冊邏輯，發送 POST 請求到後端 API
      const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        setRegisterMessage('註冊成功');
        // 在實際應用中，你可能會在這裡處理註冊成功的相應邏輯，例如導向到登入頁面
      } else {
        setRegisterMessage('註冊失敗，請檢查資料並重試');
      }
    } catch (error) {
      setRegisterMessage('發生錯誤');
    }
  };

  const handleKeyDown = event => {
    if (event.key === 'Enter') {
      handleRegister();
    }
  };

  return (
    <div>
      <h2>註冊頁面</h2>
      <div>
        <label>使用者名稱：</label>
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
      <button onClick={handleRegister}>註冊</button>
      <p>{registerMessage}</p>
    </div>
  );
}

export default Register;
