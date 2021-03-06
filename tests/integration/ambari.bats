#!/usr/bin/env bats

load ../utils/commands
load ../utils/mock_parameters

@test "change ambari pwd" {
  OUTPUT=$(change-ambari-password --name aaaaa --old-password 1234 --new-password 4321 --ambari-user admin 2>&1 | awk '{printf "%s",$0} END {print ""}' | grep -o '{.*}' | jq ' . |  [to_entries[].key] == ["oldPassword","password","userName"]')

  [[ "${OUTPUT}" ==  true ]]
}