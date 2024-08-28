#!/usr/bin/env bats

#
# Verify trapping & emulation of "ioctl" to set the immutable flag using chattr.
#

load ../../helpers/run
load ../../helpers/docker
load ../../helpers/fs
load ../../helpers/sysbox-health

function teardown() {
  sysbox_log_check
}

@test "chattr immutable flag" {
  local syscont=$(docker_run --rm ${CTR_IMG_REPO}/ubuntu:latest tail -f /dev/null)

  # Create a test file
  docker exec "$syscont" bash -c "echo 'test content' > /test_file"
  [ "$status" -eq 0 ]

  # Set the immutable flag
  docker exec "$syscont" bash -c "chattr +i /test_file"
  [ "$status" -eq 0 ]

  # Verify the immutable flag is set
  docker exec "$syscont" bash -c "lsattr /test_file | grep -q '^----i---------e----- /test_file'"
  [ "$status" -eq 0 ]

  # Try to modify the file (should fail)
  docker exec "$syscont" bash -c "echo 'new content' > /test_file"
  [ "$status" -ne 0 ]

  # Try to delete the file (should fail)
  docker exec "$syscont" bash -c "rm /test_file"
  [ "$status" -ne 0 ]

  # Remove the immutable flag
  docker exec "$syscont" bash -c "chattr -i /test_file"
  [ "$status" -eq 0 ]

  # Verify the immutable flag is removed
  docker exec "$syscont" bash -c "lsattr /test_file | grep -q '^--------------e----- /test_file'"
  [ "$status" -eq 0 ]

  # Now modification should succeed
  docker exec "$syscont" bash -c "echo 'new content' > /test_file"
  [ "$status" -eq 0 ]

  docker_stop "$syscont"
}
