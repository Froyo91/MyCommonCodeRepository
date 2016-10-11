package com.android.app.activity;

import android.content.Intent;
import android.os.Bundle;
import android.support.v4.view.PagerAdapter;
import android.support.v4.view.ViewPager;
import android.view.GestureDetector;
import android.view.MotionEvent;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;

import com.android.app.R;
import com.android.lib.annotation.Click;
import com.android.lib.annotation.Initialize;
import com.dfy.net.comment.store.UserStore;
import com.tendcloud.tenddata.TCAgent;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Administrator on 2015/6/26.
 */
public class FirstShowActivity extends AppBaseActivity implements ViewPager.OnPageChangeListener {
    @Initialize
    ViewPager viewPager;
    @Initialize
    View btn;
    @Click
    View sale;
    @Click
    View rent;
    @Click
    View publish;
    int[] srcs = new int[]{R.drawable.android1, R.drawable.android2, R.drawable.android3, R.drawable.android4};

    @Click
    TextView skip;

    private ImageView[] dots;

    // 记录当前选中位置
    private int currentIndex;

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
        findAllViewByRId(R.id.class);

        mygesture = new GestureDetector(this, new GuideViewTouch());

        TCAgent.onEvent(this, "选择模式页面");

        List<View> images = new ArrayList<>();
        for (int i = 0; i < srcs.length; i++) {
            View view = new View(this);
            view.setBackgroundResource(srcs[i]);
            images.add(view);
        }

        viewPager.setAdapter(new Adapter(images));
        viewPager.addOnPageChangeListener(this);
        btn.setVisibility(View.GONE);

        initDots();
    }

    private void initDots() {
        LinearLayout ll = (LinearLayout) findViewById(R.id.ll);
        dots = new ImageView[srcs.length];
        // 循环取得小点图片
        for (int i = 0; i < srcs.length; i++) {
            dots[i] = (ImageView) ll.getChildAt(i);
            dots[i].setBackgroundResource(R.drawable.d1);// 都设为灰
        }

        currentIndex = 0;
        dots[currentIndex].setBackgroundResource(R.drawable.d2);// 设置为白色
    }

    private void setCurrentDot(int position) {
        if (position < 0 || position > srcs.length - 1 || currentIndex == position) {
            return;
        }
        dots[position].setBackgroundResource(R.drawable.d2);
        dots[currentIndex].setBackgroundResource(R.drawable.d1);
        currentIndex = position;
    }

    @Override
    public void onClick(View v) {
        super.onClick(v);

        if (v.getId() == R.id.skip) {
            viewPager.setCurrentItem(dots.length - 1);
            return;
        }

        UserStore.setFirstStart(false);
        finish();

        if (v.getId() == R.id.sale) {
            UserStore.setBusinessType(false);
        }
        if (v.getId() == R.id.rent) {
            UserStore.setBusinessType(true);
        }
        if (v.getId() == R.id.publish) {
            TCAgent.onEvent(this, "我要发布房源");
            Intent intent = new Intent(this, MainActivityV2.class);
            Bundle bundle = new Bundle();
            bundle.putString(MainActivityV2.ACTION_PUBLISH, MainActivityV2.ACTION_PUBLISH);
            intent.putExtras(bundle);
            startActivity(intent);
        } else {
            startActivity(new Intent(this, MainActivityV2.class));
            overridePendingTransition(0, 0);
        }
    }

    private void scrollToMapView(){
        UserStore.setFirstStart(false);
        finish();

        UserStore.setBusinessType(false);

        startActivity(new Intent(this, MainActivityV2.class));
        overridePendingTransition(0, 0);
    }

    @Override
    public void onPageScrolled(int i, float v, int i1) {

    }

    @Override
    public void onPageSelected(int i) {
        btn.setVisibility((i + 1) == viewPager.getAdapter().getCount() ? View.VISIBLE : View.GONE);

        skip.setVisibility((i + 1) == viewPager.getAdapter().getCount() ? View.INVISIBLE : View.VISIBLE);

        setCurrentDot(i);
    }

    @Override
    public void onPageScrollStateChanged(int i) {

    }

    private class Adapter extends PagerAdapter {
        List<View> images;

        public Adapter(List<View> images) {
            this.images = images;
        }

        @Override
        public int getCount() {
            return images.size();
        }

        @Override
        public boolean isViewFromObject(View view, Object o) {
            return view == o;
        }

        @Override
        public void destroyItem(ViewGroup container, int position, Object object) {
            container.removeView(images.get(position));
        }

        @Override
        public Object instantiateItem(ViewGroup container, int position) {
            container.addView(images.get(position), new ViewGroup.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.MATCH_PARENT));
            return images.get(position);
        }
    }
}
