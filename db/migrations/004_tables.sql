CREATE TABLE clusters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE houses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cluster_id UUID NOT NULL REFERENCES clusters(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE entrances (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    house_id UUID NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    number INTEGER NOT NULL,
    floors_count INTEGER NOT NULL
);

CREATE TABLE floors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entrance_id UUID NOT NULL REFERENCES entrances(id) ON DELETE CASCADE,
    number INTEGER NOT NULL
);

CREATE TABLE apartments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    floor_id UUID NOT NULL REFERENCES floors(id) ON DELETE CASCADE,
    number TEXT NOT NULL
);

CREATE TABLE equipment_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL
);

CREATE TABLE floor_equipment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    floor_id UUID NOT NULL REFERENCES floors(id) ON DELETE CASCADE,
    equipment_type_id UUID NOT NULL REFERENCES equipment_types(id),
    installed BOOLEAN DEFAULT false,
    connected BOOLEAN DEFAULT false
);

CREATE TABLE work_types (
    code TEXT PRIMARY KEY,
    system_code TEXT NOT NULL,
    scope TEXT NOT NULL,
    name TEXT NOT NULL
);


CREATE TABLE basement_works (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    house_id UUID NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    work_type_id UUID NOT NULL REFERENCES work_types(id),
    completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMPTZ
);

CREATE TABLE floor_works (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    floor_id UUID NOT NULL REFERENCES floors(id) ON DELETE CASCADE,
    work_type_id UUID NOT NULL REFERENCES work_types(id),
    completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMPTZ
);

CREATE TABLE apartment_works (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    apartment_id UUID NOT NULL REFERENCES apartments(id) ON DELETE CASCADE,
    work_type_id UUID NOT NULL REFERENCES work_types(id),
    completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMPTZ
);
