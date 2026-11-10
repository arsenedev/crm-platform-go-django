
CREATE TABLE audit_logs (
   id UUID PRIMARY KEY,
   user_id UUID,
   action VARCHAR(255),
   resource VARCHAR(255),
   resource_id VARCHAR(255),
   created_at TIMESTAMP DEFAULT NOW()
);