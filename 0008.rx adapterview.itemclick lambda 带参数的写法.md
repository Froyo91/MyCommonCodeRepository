RxAdapterView.itemClicks(listview).subscribe((Integer position) -> {
    Intent intent = new Intent(this, HouseDetailActivity.class);
    intent.putExtra("id", position);
    startActivity(intent);
});

