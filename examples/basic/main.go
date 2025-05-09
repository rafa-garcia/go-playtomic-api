// Example showing basic usage of the Playtomic API client
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rafa-garcia/go-playtomic-api/client"
	"github.com/rafa-garcia/go-playtomic-api/models"
)

func main() {
	// Create a client with options
	c := client.NewClient(
		client.WithTimeout(15*time.Second),
		client.WithRetries(2),
	)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Search for classes
	fmt.Println("Searching for classes...")
	classes, err := searchClasses(ctx, c)
	if err != nil {
		log.Fatalf("Error searching classes: %v", err)
	}

	// Display classes
	fmt.Printf("Found %d classes\n", len(classes))
	for i, class := range classes {
		if i >= 3 {
			fmt.Println("...")
			break
		}

		fmt.Printf("- %s at %s (%s to %s)\n",
			getClassTitle(class),
			class.Tenant.TenantName,
			class.StartDate,
			class.EndDate)
	}

	fmt.Println()

	// Search for matches
	fmt.Println("Searching for matches...")
	matches, err := searchMatches(ctx, c)
	if err != nil {
		log.Fatalf("Error searching matches: %v", err)
	}

	// Display matches
	fmt.Printf("Found %d matches\n", len(matches))
	for i, match := range matches {
		if i >= 3 {
			fmt.Println("...")
			break
		}

		fmt.Printf("- %s match at %s (%s): %d of %d players\n",
			match.MatchType,
			match.Tenant.TenantName,
			match.StartDate,
			countPlayers(match),
			totalPlayerSlots(match))
	}

	fmt.Println()

	// Search for lessons
	fmt.Println("Searching for lessons...")
	lessons, err := searchLessons(ctx, c)
	if err != nil {
		log.Fatalf("Error searching lessons: %v", err)
	}

	// Display lessons
	fmt.Printf("Found %d lessons\n", len(lessons))
	for i, lesson := range lessons {
		if i >= 3 {
			fmt.Println("...")
			break
		}

		fmt.Printf("- %s at %s (%s): %d of %d players, %d available spots\n",
			lesson.TournamentName,
			lesson.Tenant.TenantName,
			lesson.StartDate,
			len(lesson.RegisteredPlayers),
			lesson.MaxPlayers,
			lesson.AvailablePlaces)

		// Demonstrate model conversion if there are players
		if len(lesson.RegisteredPlayers) > 0 {
			// Convert lesson player to standard player
			lessonPlayer := &lesson.RegisteredPlayers[0]
			player := models.LessonPlayerToPlayer(lessonPlayer)

			fmt.Printf("  Player: %s (converted from LessonPlayer)\n", player.Name)
		}
	}
}

// searchClasses demonstrates searching for classes
func searchClasses(ctx context.Context, c *client.Client) ([]models.Class, error) {
	// Build search parameters
	classParams := &models.SearchClassesParams{
		Sort:             "start_date,ASC",
		Status:           "PENDING,IN_PROGRESS",
		Type:             "COURSE,PUBLIC",
		IncludeSummary:   true,
		Size:             100,
		Page:             0,
		CourseVisibility: "PUBLIC",
		FromStartDate:    time.Now().Format("2006-01-02") + "T00:00:00",
	}

	// Add tenant IDs if provided
	tenantID := os.Getenv("PLAYTOMIC_TENANT_ID")
	if tenantID != "" {
		classParams.TenantIDs = []string{tenantID}
	}

	return c.GetClasses(ctx, classParams)
}

// searchMatches demonstrates searching for matches
func searchMatches(ctx context.Context, c *client.Client) ([]models.Match, error) {
	// Build search parameters
	matchParams := &models.SearchMatchesParams{
		Sort:          "start_date,DESC",
		HasPlayers:    true,
		SportID:       "PADEL",
		Visibility:    "VISIBLE",
		FromStartDate: time.Now().Format("2006-01-02") + "T00:00:00",
		Size:          100,
		Page:          0,
	}

	// Add tenant IDs if provided
	tenantID := os.Getenv("PLAYTOMIC_TENANT_ID")
	if tenantID != "" {
		matchParams.TenantIDs = []string{tenantID}
	}

	return c.GetMatches(ctx, matchParams)
}

// searchLessons demonstrates searching for lessons
func searchLessons(ctx context.Context, c *client.Client) ([]models.Lesson, error) {
	// Build search parameters
	lessonParams := &models.SearchLessonsParams{
		Sort:                 "start_date,ASC",
		Status:               "REGISTRATION_OPEN,REGISTRATION_CLOSED,IN_PROGRESS",
		TournamentVisibility: "PUBLIC",
		Size:                 100,
		Page:                 0,
		FromStartDate:        time.Now().Format("2006-01-02") + "T00:00:00",
	}

	// Add tenant ID if provided
	// Note: Lessons API only accepts a single tenant_id, not a list
	tenantID := os.Getenv("PLAYTOMIC_TENANT_ID")
	if tenantID != "" {
		lessonParams.TenantID = tenantID
	}

	return c.GetLessons(ctx, lessonParams)
}

// Helper function to get class title
func getClassTitle(class models.Class) string {
	if class.CourseSummary != nil && class.CourseSummary.Name != "" {
		return class.CourseSummary.Name
	}
	return class.Resource.Name
}

// Helper function to count registered players in a match
func countPlayers(match models.Match) int {
	count := 0
	for _, team := range match.Teams {
		count += len(team.Players)
	}
	return count
}

// Helper function to calculate total player slots in a match
func totalPlayerSlots(match models.Match) int {
	return match.MinPlayersPerTeam * len(match.Teams)
}
