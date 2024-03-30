import type { Knex } from "knex";
import dotenv from "dotenv";

dotenv.config();
console.log(process.env);
// Update with your config settings.
const config: { [key: string]: Knex.Config } = {
  development: {
    client: process.env.DB_ENGINE,
    connection: {
      database: process.env.POSTGRES_NAME,
      user: process.env.POSTGRES_USER,
      password: process.env.POSTGRES_PASSWORD,
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      tableName: "knex_migrations",
    },
  },

  staging: {
    client: process.env.DB_ENGINE,
    connection: {
      database: process.env.POSTGRES_NAME,
      user: process.env.POSTGRES_USER,
      password: process.env.POSTGRES_PASSWORD,
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      tableName: "knex_migrations",
    },
  },

  production: {
    client: "postgresql",
    connection: {
      database: process.env.POSTGRES_NAME,
      user: process.env.POSTGRES_USER,
      password: process.env.POSTGRES_PASSWORD,
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      tableName: "knex_migrations",
    },
  },
};

console.log(config);
module.exports = config;
