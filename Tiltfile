## FRONTEND APP
k8s_yaml('k8s/tilt-depl/apps/frontend.yaml')
sync_src = sync('./app/frontend/app', '/code/app')
docker_build('frontend-tilt', 'app/frontend', live_update=[sync_src], entrypoint='flask --app app.main run --debug', build_args={'ENV': 'DEV_TILT'})
k8s_resource('frontend-v1', port_forwards='3000:5000')

## DETAILS APP
k8s_yaml('k8s/tilt-depl/apps/details.yaml')
docker_build('details-tilt', 'app/details')
k8s_resource('details-v1', port_forwards='3001:7777')

## PAYMENT APP
k8s_yaml('k8s/tilt-depl/apps/payment.yaml', allow_duplicates=True)
docker_build('payment-tilt-v1', 'app/payment')
k8s_resource('payment-v1', port_forwards='3002:8888')

## REVIEWS APP
k8s_yaml('k8s/tilt-depl/apps/reviews.yaml', allow_duplicates=True)
docker_build('reviews-tilt-v1', 'app/reviews')
k8s_resource('reviews-tilt-v1-depl', port_forwards='3003:9999')