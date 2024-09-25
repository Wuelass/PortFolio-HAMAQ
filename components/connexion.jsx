import React from 'react';

const connexion = () => {
  return (
    <div style={{ maxWidth: '(300px', margin: '10 auto' }}>
      <form>
        <div>
          <label htmlFor="username">Username</label>
          <input type="text" id="username" name="username" />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input type="text" id="password" name="password" />
        </div>
        <button type="submit">Envoyer</button>
      </form>
    </div>
  );
};

export default connexion;
