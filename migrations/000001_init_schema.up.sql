CREATE TABLE IF NOT EXISTS tracks (
                                      id UUID PRIMARY KEY,
                                      track_number TEXT NOT NULL,
                                      carrier TEXT NOT NULL,
                                      status TEXT NOT NULL,
                                      last_update_at TIMESTAMPTZ NOT NULL,
                                      created_at TIMESTAMPTZ NOT NULL,
                                      UNIQUE(track_number, carrier)
    );

CREATE INDEX IF NOT EXISTS idx_tracks_last_update ON tracks(last_update_at);