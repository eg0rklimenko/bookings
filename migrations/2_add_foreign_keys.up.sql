ALTER TABLE room_restrictions
  ADD CONSTRAINT reservation_id_fk FOREIGN KEY (reservation_id) REFERENCES reservations(id);

ALTER TABLE room_restrictions
  ADD CONSTRAINT room_id_fk FOREIGN KEY (room_id) REFERENCES rooms(id);

ALTER TABLE reservations
  ADD CONSTRAINT room_id_fk FOREIGN KEY (room_id) REFERENCES rooms(id);
