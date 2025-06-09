**Title:**
Duplicate "Permissions" Fields Displayed in CMS Since Release 6.20.7

**Description:**
After upgrading to Release 6.20.7, the CMS interface shows two separate fields labeled "Permissions" where only one should appear. This is causing confusion among users when setting or reviewing content permissions.

**Steps to Reproduce:**

1. Log in to the CMS.
2. Navigate to any content item that has configurable permissions.
3. Observe the "Permissions" section — two identical fields are present.

**Expected vs Actual Behavior:**

* **Expected:**
  Only one "Permissions" field should be displayed for managing content permissions.

* **Actual:**
  Two identical "Permissions" fields are displayed, leading to ambiguity.

**Environment (if known):**

* CMS Version: 6.20.7
* Affected Browsers: All major browsers (Chrome, Firefox, Edge)
* Affected Users: All CMS users

**Severity or Impact:**
**Medium** — This issue is causing user confusion but does not block core functionality. However, there is a risk of incorrect permission settings due to ambiguity.
