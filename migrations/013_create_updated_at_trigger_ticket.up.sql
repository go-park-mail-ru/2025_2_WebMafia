CREATE TRIGGER update_ticket_updated_at
    BEFORE UPDATE ON "ticket"
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();