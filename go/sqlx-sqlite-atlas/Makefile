
ATLAS=atlas

all:
	@echo "No default action."

demo.db: schema.sql
	$(ATLAS) schema apply \
	  --auto-approve \
	  -u "sqlite://$@" \
	  --to "file://$?" \
	  --dev-url "sqlite://dev?mode=memory"

clean:
	-rm -f demo.db
