package adapter_test

// func Test_groupProductPg_DeleteGroupProduct(t *testing.T) {
// 	if !onlinetest {
// 		t.Skip("skipping online tests")
// 	}

// 	ctx := context.Background()

// 	beginner := txcom.NewGormTxBeginner(gdb)
// 	repo := adapter.NewGroupProductPg(gdb)

// 	tx, err := beginner.Begin()
// 	assert.Nil(t, err)
// 	defer tx.Rollback()

// 	affected, err := repo.DeleteGroupProduct(ctx, tx, group2.ID, product2.ID)
// 	assert.Nil(t, err)
// 	assert.EqualValues(t, 1, affected)

// 	assert.Nil(t, tx.Commit())
// }
