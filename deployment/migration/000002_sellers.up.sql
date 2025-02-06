
CREATE TABLE IF NOT EXISTS sellers (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    phone TEXT,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by_id BIGINT,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by_id BIGINT,
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE NULLS NOT DISTINCT (tenant_id, email, deleted_at),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);

-- Index on tenant_id to speed up lookups on a specific tenant
CREATE INDEX idx_sellers_tenant_id ON sellers(tenant_id);

-- Index on email for faster lookups
CREATE INDEX idx_sellers_email ON sellers(email);

-- Composite index on tenant_id and email for optimized lookups
CREATE INDEX idx_sellers_tenant_email ON sellers(tenant_id, email);



