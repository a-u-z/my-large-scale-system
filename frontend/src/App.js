// App.js
import React, { useState } from 'react';
import Login from './Login';
import Register from './RegistrationPage'; // 假設你還需要一個註冊組件

function App() {
  const [isLoginVisible, setIsLoginVisible] = useState(true);

  const togglePage = () => {
    setIsLoginVisible(!isLoginVisible);
  };

  return (
    <div>
      {/* 切換頁面的按鈕 */}
      <button onClick={togglePage}>
        {isLoginVisible ? '切換到註冊頁面' : '切換到登入頁面'}
      </button>

      {/* 根據狀態顯示登入或註冊頁面 */}
      {isLoginVisible ? <Login /> : <Register />}
    </div>
  );
}

export default App;
