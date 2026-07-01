package dog

import "testing"

func TestFindByNameIsCaseInsensitive(t *testing.T) {
	repo := NewRepository()

	got, ok := repo.FindByName("mochi")
	if !ok {
		t.Fatal("expected to find Mochi")
	}
	if got.Name != "Mochi" {
		t.Fatalf("expected Mochi, got %q", got.Name)
	}
}

func TestAllReturnsCopy(t *testing.T) {
	repo := NewRepository()

	dogs := repo.All()
	dogs[0].Name = "changed"

	fresh := repo.All()
	if fresh[0].Name == "changed" {
		t.Fatal("All should return a copy")
	}
}
