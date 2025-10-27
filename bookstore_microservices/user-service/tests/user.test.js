const request = require('supertest');
const app = require('../index');
const fs = require('fs');

describe('User Service basic flows', () => {
  beforeAll(() => {
    // reset DB file
    fs.writeFileSync(__dirname + '/../storage/users.json', '[]');
  });

  test('signup and get profile', async () => {
    const signup = await request(app).post('/auth/signup').send({ username: 'alice', email: 'alice@example.com', password: 'pass' });
    expect(signup.statusCode).toBe(201);
    const id = signup.body.id;
    const get = await request(app).get('/users/' + id);
    expect(get.statusCode).toBe(200);
    expect(get.body.username).toBe('alice');
    expect(get.body.email).toBe('alice@example.com');
    expect(get.body).not.toHaveProperty('password');
  });

  test('login succeeds', async () => {
    const login = await request(app).post('/auth/login').send({ email: 'alice@example.com', password: 'pass' });
    expect(login.statusCode).toBe(200);
    expect(login.body.message).toBe('login successful');
  });
});
