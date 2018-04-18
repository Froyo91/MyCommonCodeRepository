在onScroll方法中添加：
View c = listView.getChildAt(0);
if (c == null) {
    return 0;
}

int firstVisiblePosition = listView.getFirstVisiblePosition();
int top = c.getTop();

int headerHeight = 0;
if (firstVisiblePosition >= 1) {
    headerHeight = listView.getHeight();
}
return -top + firstVisiblePosition * c.getHeight() + headerHeight;
