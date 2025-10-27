// Simple in-memory + file persistence DB for demo (no real MySQL connection here).
// In a full implementation this would connect to MySQL using mysql2.

const fs = require('fs');
const path = require('path');
const file = path.join(__dirname, 'users.json');

let users = [];
try {
  users = JSON.parse(fs.readFileSync(file));
} catch (e) {
  users = [];
}

function save() {
  fs.writeFileSync(file, JSON.stringify(users, null, 2));
}

function create({ username, email, password }) {
  const id = users.length ? users[users.length-1].id + 1 : 1;
  const user = { id, username, email, password };
  users.push(user);
  save();
  return user;
}

function findByEmail(email) {
  return users.find(u => u.email === email);
}

function findById(id) {
  return users.find(u => u.id === id);
}

module.exports = { create, findByEmail, findById };
