ALTER TABLE room_restrictions
  DROP CONSTRAINT reservation_id_fk;

ALTER TABLE room_restrictions
  DROP CONSTRAINT room_id_fk;

ALTER TABLE reservations
  DROP CONSTRAINT room_id_fk;
