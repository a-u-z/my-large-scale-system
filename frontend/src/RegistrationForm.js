// RegistrationForm.js
import React, { useState } from 'react';

function RegistrationForm() {
  const [formData, setFormData] = useState({
    // Define form fields and their initial values
    username: '',
    email: '',
    password: '',
  });

  const handleInputChange = e => {
    // Handle form input changes
    const { name, value } = e.target;
    setFormData(prevData => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = e => {
    // Handle form submission
    e.preventDefault();
    // Perform registration logic with formData
  };

  return (
    <form onSubmit={handleSubmit}>
      {/* Render form fields and handle changes */}
      <label>
        Username:
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleInputChange}
        />
      </label>
      {/* Other form fields */}
      <button type="submit">Register</button>
    </form>
  );
}

export default RegistrationForm;
