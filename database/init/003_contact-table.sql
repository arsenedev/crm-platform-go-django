CREATE TABLE contacts (
   id UUID PRIMARY KEY,
   customer_id UUID REFERENCES customers(id),
   firstname VARCHAR(255),
   lastname VARCHAR(255),
   email VARCHAR(255),
   phone VARCHAR(255),
   created_at TIMESTAMP DEFAULT NOW()
);