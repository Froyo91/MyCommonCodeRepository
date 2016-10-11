    private GestureDetector mygesture;
    private static final int FLING_MIN_DISTANCE = 180;
    private static final int FLING_MIN_VELOCITY = 300;

    private class GuideViewTouch extends GestureDetector.SimpleOnGestureListener {
        @Override
        public boolean onFling(MotionEvent e1, MotionEvent e2, float velocityX, float velocityY) {
			if (currentIndex == (srcs.length - 1) && e1.getX() - e2.getX() > FLING_MIN_DISTANCE && Math.abs(e1.getY() - e2.getY()) < 200 && Math.abs(velocityX) > FLING_MIN_VELOCITY) {
                scrollToMapView();
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

        mygesture = new GestureDetector(this, new GuideViewTouch());
    }
