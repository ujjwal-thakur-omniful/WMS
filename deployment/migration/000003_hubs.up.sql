CREATE TABLE IF NOT EXISTS hubs (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    tenant_id BIGINT NOT NULL,
    location POINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);

-- Index on tenant_id to speed up lookups for a specific tenant
CREATE INDEX idx_hubs_tenant_id ON hubs(tenant_id);

-- Index on name for faster searches by name
CREATE INDEX idx_hubs_name ON hubs(name);