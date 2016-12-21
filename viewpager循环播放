private final class MyDemo implements OnPageChangeListener {

    private int mCurrentPosition;

    @Override
    public void onPageScrollStateChanged(int state) {
        if (state == ViewPager.SCROLL_STATE_IDLE) {
            if (mCurrentPosition == viewPager.getAdapter().getCount() - 1) {
                viewPager.setCurrentItem(1, false);
            }
            else if (mCurrentPosition == 0) {
                viewPager.setCurrentItem(viewPager.getAdapter().getCount() - 2, false);
            }
        }
    }

    @Override
    public void onPageScrolled(int scrolledPosition, float percent, int pixels) {
        
    }

    @Override
    public void onPageSelected(int position) {
        mCurrentPosition = position;
    }

}

在这里，还有一点没有写出来，就是在原有数据的前面加一条数据（最后一条数据），在最后面加一条数据（第一条数据）。那么现在数据的总量就是原有数据+2。上面是我自己的代码，可能还需要视实际情况进行修改。

