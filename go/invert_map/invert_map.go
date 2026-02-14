// Inverting the Map (Medium)
// Goal: Value manipulation.

// Create a map where keys are usernames (string) and values are team names (string).

// Create a new map that inverts this: the keys should be team names and the values
// should be a slice of usernames belonging to that team.

package main

func invertMap(username_team map[string]string) map[string][]string {
	team_usernames := make(map[string][]string, len(username_team))

	for username, team := range username_team {
		team_usernames[team] = append(team_usernames[team], username)
	}

	return team_usernames
}
