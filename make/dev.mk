ETC_HOSTS=/etc/hosts

APP_NAMESPACE ?= $(NAMESPACE)
NAMESPACE ?= "alexeykazakov-dev"

.PHONY: create-namespace
## Create the test namespace
create-namespace:
	$(Q)-echo "Creating Namespace"
	$(Q)-oc new-project $(NAMESPACE)
	$(Q)-echo "Switching to the namespace $(NAMESPACE)"
	$(Q)-oc project $(NAMESPACE)

.PHONY: use-namespace
## Log in as system:admin and enter the test namespace
use-namespace:
	$(Q)-echo "Using to the namespace $(NAMESPACE)"
	$(Q)-oc project $(NAMESPACE)

.PHONY: clean-namespace
## Delete the test namespace
clean-namespace:
	$(Q)-echo "Deleting Namespace"
	$(Q)-oc delete project $(NAMESPACE)

.PHONY: reset-namespace
## Delete an create the test namespace and deploy rbac there
reset-namespace: clean-namespace create-namespace

.PHONY: deploy
## Deploy htmlserver
deploy: create-namespace build docker-image docker-push apply-resources print-route

.PHONY: apply-resources
## Apply htmlserver resources
apply-resources:
	$(Q)-oc process -f ./deploy/htmlserver.yaml \
        -p NAMESPACE=${NAMESPACE} \
        -p IMAGE=${IMAGE} \
        | oc apply -f -

.PHONY: print-route
print-route:
	@echo "------------------------------------------------------------------"
	@echo "Deployment complete! Waiting for the htmlserver service route."
	@echo -n "."
	@while [[ -z `oc get routes htmlserver -n ${NAMESPACE} 2>/dev/null` ]]; do \
		if [[ $${NEXT_WAIT_TIME} -eq 100 ]]; then \
            echo ""; \
            echo "The timeout of waiting for the service route has been reached. Try to run 'make print-route' later or check the deployment logs"; \
            exit 1; \
		fi; \
		echo -n "."; \
		sleep 1; \
	done
	@echo ""
	$(eval ROUTE = $(shell oc get routes htmlserver -n ${NAMESPACE} -o=jsonpath='{.spec.host}'))
	@echo Access the Landing Page here: https://${ROUTE}
	@echo "------------------------------------------------------------------"