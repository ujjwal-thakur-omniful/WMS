CREATE TABLE IF NOT EXISTS inventories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    hub_id BIGINT NOT NULL,
    sku_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (hub_id) REFERENCES hubs(id) ON DELETE CASCADE,
    FOREIGN KEY (sku_id) REFERENCES sku(id) ON DELETE CASCADE
);

-- Index on hub_id to speed up lookups for a specific hub
CREATE INDEX idx_inventories_hub_id ON inventories(hub_id);

-- Index on sku_id to speed up lookups for a specific SKU
CREATE INDEX idx_inventories_sku_id ON inventories(sku_id);

-- Composite index on (hub_id, sku_id) to optimize queries that filter by both
CREATE INDEX idx_inventories_hub_sku ON inventories(hub_id, sku_id);
