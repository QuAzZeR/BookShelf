import type { Knex } from 'knex';

export async function up(knex: Knex): Promise<void> {
  knex.schema.createTable('book', (table) => {
    table.uuid('id').primary();
    table.string('name').notNullable();
    table.string('eng_name');
    table.string('author').notNullable();
    table.string('transalator').notNullable();
    table.string('cover_url').notNullable();
    table.string('language');
    table.string('genre').notNullable();
    table.datetime('created_at');
    table.datetime('updated_at');
  });
}

export async function down(knex: Knex): Promise<void> {
  knex.schema.dropTableIfExists('book');
}
