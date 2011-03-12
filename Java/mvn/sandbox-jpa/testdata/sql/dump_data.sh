#
# The following is a script used to generate the seed data.  Use it
# as a template for generating future datasets, it is NOT intended to
# be blindly run.
#
mysqldump -u sandbox -pdeveloper junit_sandbox_jpa \
  --skip-add-locks \
  --skip-disable-keys \
  --skip-extended-insert \
  --no-create-db \
  --no-create-info \
  bags boxes extras
