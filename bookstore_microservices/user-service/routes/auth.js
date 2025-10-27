const express = require('express');
const router = express.Router();
const usersDB = require('../storage/users_db');

// Signup - simple (passwords stored in plaintext for demo only)
router.post('/signup', (req, res) => {
  const { username, email, password } = req.body;
  if(!username || !email || !password) return res.status(400).json({error:'missing fields'});
  const user = usersDB.create({ username, email, password });
  res.status(201).json({ id: user.id, username: user.username, email: user.email });
});

// Login - simple check
router.post('/login', (req, res) => {
  const { email, password } = req.body;
  const user = usersDB.findByEmail(email);
  if(!user || user.password !== password) return res.status(401).json({error:'invalid credentials'});
  res.json({ message: 'login successful', id: user.id, username: user.username });
});

module.exports = router;
