package ds

import (
	"context"
	"testing"
)
func IniciarDependencias() (context.Context,*fakeRepository, Repository){
	ctx := context.Background()
	fr := NewFakeRepository()
	s:= NewRepository(fr.db)
	return ctx, fr, s
}

func TestRepository_GetAll(t *testing.T) {
	ctx := context.Background()


}