在Activity中：
@Override
public void onBackPressed() {
    int count = getSupportFragmentManager().getBackStackEntryCount();//使用的manager根据实际情况

    if (count == 0) {
        exitConfirm();
    } else {
        getSupportFragmentManager().popBackStack();
    }
}
