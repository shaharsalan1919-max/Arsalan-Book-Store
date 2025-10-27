const express = require('express');
const router = express.Router();
const usersDB = require('../storage/users_db');

router.get('/:id', (req, res) => {
  const id = parseInt(req.params.id);
  const user = usersDB.findById(id);
  if(!user) return res.status(404).json({error:'user not found'});
  // do not send password
  const { password, ...safe } = user;
  res.json(safe);
});

module.exports = router;
