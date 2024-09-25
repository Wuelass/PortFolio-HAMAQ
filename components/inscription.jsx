import React from 'react';

const inscription = () => {
  return (
    <div style={{ maxWidth: '(500px', margin: '0 auto' }}>
      <form>
        <div>
          <label htmlFor="username">Username</label>
          <input type="text" id="username" name="username" />
        </div>
        <div>
          <label htmlFor="email">Email</label>
          <input type="text" id="email" name="email" />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input type="text" id="password" name="password" />
        </div>
        <div>
          <label htmlFor="confirm password">Confirm password</label>
          <input type="confirm password" id="confirm password" name="confirm password" />
        </div>
        <button type="submit">Envoyer</button>
      </form>
    </div>
  );
};

export default inscription;
