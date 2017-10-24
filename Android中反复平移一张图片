private RelativeLayout lottery_area;
private boolean toRight = true;
private boolean isRunning = false;

final ImageView folder = list.findViewById(R.id.fold);
lottery_area = list.findViewById(R.id.lottery_area);
folder.setOnClickListener(new View.OnClickListener() {
    @Override
    public void onClick(View view) {
        if(!isRunning) {
            isRunning = true;
            final float moveX;
            if(toRight){
                toRight = false;
                moveX = utils.Utils.getWindowDisplay(getActivity()).widthPixels - utils.Utils.dip2px(getActivity().getApplicationContext(), 60);
            } else {
                toRight = true;
                moveX = - utils.Utils.getWindowDisplay(getActivity()).widthPixels + utils.Utils.dip2px(getActivity().getApplicationContext(), 60);
            }
            //初始化
            Animation translateAnimation = new TranslateAnimation(0, moveX, 0, 0);
            translateAnimation.setDuration(1000);
            lottery_area.startAnimation(translateAnimation);
            translateAnimation.setAnimationListener(new Animation.AnimationListener() {
                @Override
                public void onAnimationStart(Animation animation) {

                }

                @Override
                public void onAnimationEnd(Animation animation) {
                    isRunning = false;

                    int left = lottery_area.getLeft()+(int)moveX;
                    int top = lottery_area.getTop();
                    int width = lottery_area.getWidth();
                    int height = lottery_area.getHeight();
                    lottery_area.clearAnimation();
                    lottery_area.layout(left, top, left+width, top+height);

                    if(toRight){
                        folder.setImageResource(R.mipmap.fold);
                    } else {
                        folder.setImageResource(R.mipmap.unfold);
                    }
                }

                @Override
                public void onAnimationRepeat(Animation animation) {

                }
            });
        }
    }
});
