create table if not exists model (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    repo_id text not null,
    filename text not null,
    input_token_price int not null,
    output_token_price int not null,
    updated_at timestamptz not null
);
