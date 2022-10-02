db.createUser({
  user: "root",
  pwd: "password",
  roles: [
    {
      role: "readWrite",
      db: "supergo",
    },
  ],
});

db.users.createIndex({ user_id: 1 }, { unique: true });
