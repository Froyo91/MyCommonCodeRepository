// Android welcone activity : slide left to turn to next activity in the last page of the viewpager
// 在使用viewpager的引导页的最后一页上，向左滑动切换到下一个界面的收拾代码片段

private GestureDetector mygesture;
private static final int FLING_MIN_DISTANCE = 180;
private static final int FLING_MIN_VELOCITY = 300;

private class GuideViewTouch extends GestureDetector.SimpleOnGestureListener {
	@Override
	public boolean onFling(MotionEvent e1, MotionEvent e2, float velocityX, float velocityY) {
		if (currentIndex == (srcs.length - 1) && e1.getX() - e2.getX() > FLING_MIN_DISTANCE && Math.abs(e1.getY() - e2.getY()) < 200 && Math.abs(velocityX) > FLING_MIN_VELOCITY) {
			//currentIndex是当前viewpager所在的索引，比如viewpager共有4页，那么currentIndex取值范围是0到3，最后一页的currentIndex值就是3.
			scrollToMapView();//跳转到下一个界面的函数（自己定义的）
		}
		return false;
	}
}

@Override
public boolean dispatchTouchEvent(MotionEvent ev) {
if (mygesture != null && mygesture.onTouchEvent(ev)) {
    ev.setAction(MotionEvent.ACTION_CANCEL);
}
return super.dispatchTouchEvent(ev);
}

@Override
protected void onCreate(Bundle savedInstanceState) {
super.onCreate(savedInstanceState);
setContentView(R.layout.activity_first);

mygesture = new GestureDetector(this, new GuideViewTouch());//一般这一行都放在onCreate里初始化
}
