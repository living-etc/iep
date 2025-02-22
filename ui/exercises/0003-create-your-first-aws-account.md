When you create your first AWS account, it can be quite daunting figuring out how to do it right so your project gets off to the best start it can. There's a lot of moving parts involved and this tutorial brings them all together into a single comprehensive guide.

# Prerequisites

- A computer with an internet connection
- A valid credit or debit card. You won't be spending anything today, but you can't create an AWS account with a valid payment method.

# Walkthrough

## Part 1 - Creating the AWS account

1. Follow [this link](https://signin.aws.amazon.com/signup?request_type=register) to the AWS signup page.
2. Fill in the `Root user email address` and `AWS account name`.
3. Click `Verify email address`. This will cause an email to be sent to the `Root user email address` containing a six digit verification code. Type or paste that code into the `Verification code` box that should now be on your screen.
4. Now choose a secure password for your root user.
5. On the next page:
   - Choose `Personal - for your own projects`
   - Fill in your name, phone number and address
   - Check the box saying you accept the terms of the AWS Customer Agreement
   - Click `Continue`.
6. On the next page enter your payment card details. You'll only be charged a $1 verification amount, which will then be refunded. After that, you'll only be charged for whatever computer or storage resources you use. Click `Verify and Continue` to make the payment.
7. Once your payment card has been verified, you now need to verify your phone number. Do that on this page and click `Send SMS`.
8. You'll receive a SMS message with a 4-digit code. Enter that on the next page.
9. On the next page, choose `Basic support - Free` and click `Complete sign up`.
10. Congratulation, you now have an AWS account.
11. Click on `Go to the AWS Management Console` and we'll proceed to the next step, which is securing the account's root user.

## Part 2 - Securing the root user

1. Every AWS account has a root user that is created when the account is. This is the superuser with unlimited permissions, including being the only user that is able to close the account.
2. Before we continue, we will enable Multi-Factor Authentication on the root user. Start by clicking the account name in the upper right corner, followed by `Security Credentials`.
3. Next you want to click `Assign MFA`.
   - On the next page:
   - Give your MFA device a name, I chose `root`
   - Choose `Authenticator app` for your device type. This is the most common option.
   - Click `Next`
4. If you already have an authenticator app on your phone, you can skip step 1 on the next page, otherwise follow the link to `See a list of compatible applications` and set one up.
5. For step 2, click `Show QR code` and scan it with your authenticator app. For step 3, wait for the code to refresh itself twice, entering each code into the fields marked `MFA Code 1` and `MFA Code 2`. Click `Add MFA`.
6. You'll be brought back to `Security credentials` where you'll see a green banner confirming setup of the MFA device. You'll need to use this device every time you log in as the root user, and if you lose it you'll need to [get on a phone call](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_lost-or-broken.html#root-mfa-lost-or-broken) with AWS to reset it.
7. Scrolling down this page, you'll see several different types of credentials: `Access keys`, `CloudFront key pairs` and `X.509 Signing certificates`. There shouldn't be any credentials on the root user, now or ever. If you've just created your account, then this should already be the case.

> Never let the root user have any credentials!
>
> If you're using an existing account where any credentials have been added to the root user, I strongly suggest you delete them from the root user and recreate them on IAM users, and in the next part we'll create an IAM user for day to day use.

## Part 3 - Setting up AWS Organization

1. Using the search bar at the top of the page, search for and select AWS Organizations.
2. Choose `Create an organization`.
3. Once again, the green banner signifies success.

## Part 4 - Setting up IAM Identity Center

1. Since we're not going to be using the root user for any work, we need to create another user for this purpose.
2. In the search bar at the top of the AWS console, type `iam` and click on the `IAM Identity Center` link that appears.

> IAM vs IAM Identity Center
>
> When you search for `iam` in the search box, you'll see two services that begin with that name. `IAM` is the old way of doing this, while `IAM Identity Center` is newer and recommended when creating new accounts. This tutorial will use `IAM Identity Center`.

3. First, we need to enable `IAM Identity Center` in our account by clicking on `Enable`.
1. You'll now end up on the homepage for `IAM Identity Centre` with a green banner confirming it's set up.
1. We now need to give a name to our identity center instance.
1. Name it whatever you want and click `Save changes`.
1. You'll now see confirmation that it's renamed.

## Part 5 - Creating an IAM user

1. Now that we've set up the `IAM Identity Center`, we can create our personal administrator user.
2. Click on `Users` from the left of `IAM Identity Center` homepage.
3. Click `Add user` from the upper right.
4. On the next page, do the following:
   - Give your user a username
   - Choose to send an email with password setup instructions
   - Give the user an email address. I'm using Gmail's `+` feature to save
     creating a second Google account for this tutorial.
   - Set the first, last and display names.
   - Skip the rest of the tabs and click `Next`.
5. Choose to send an email with pNext, choose `Create group`. This will open a new tab. Ensure you keep the current tab since we'll be returning to it once we've created the group.
6. It's good practice to not attach policies directly to users, rather it's best to give groups permissions, and add users to that group. We're going to create an `administrators` group.
7. Fill in the name and description before clicking `Create group`.
8. We'll see the now familiar green banner saying we've successfully created our group.
9. Close the group creation tab and return to the user creation tab. Click the refresh button in the upper right to cause our `administrators` group to appear. Select it and click `Next`.
10. Confirm the details are correct before clicking `Add user`.
11. Once again, we'll see a green banner telling us the user was created successfully.
12. Finally we need to give our new user some permissions. It's always a good idea to attach permissions to groups, and add users to those groups, instead of giving permissions directly to users, so that's what we'll do here.
13. From the IAM Identity Center sidebar, choose `Permission Sets` followed by `Create permission set`.
14. Next, choose `Administrator Access` from the predefined permission sets followed by `Next`.
15. Leave the name as it is and give it a description if you want, before clicking `Next`.
16. Review the settings before clicking `Create`.
17. Now we've created the permission set, we need to associate it with an account. From the sidebar again choose `AWS accounts`, check the box next to the account name, and click `Assign users or groups`.
18. Choose the administrators group we created earlier and click `Next`.
19. Select the permission set we created earlier and click `Next`.
20. Review your changes and click `Submit`.
21. The IAM user is now created with administrator permissions, and before we start using it we need to set up and secure the user.

## Part 6 - Setting up and securing the IAM user

1. When we created the user in the end of last section, an email was sent to the email address we provided inviting our user to IAM Identity Center. Go to your email, find this message and click the link that says `Accept invitation`.
2. Choose a password, enter the confirmation and click `Set new password`.
3. Sign in with your username and the password you just created.
4. You'll now be prompted to set up MFA for your IAM user.

> Different MFA devices
>
> This MFA device is separate from the one you set up for the root user. You can use the same device to manage both accounts, but you need to add your IAM user account separately from the root user.

5. Choose `Authenticator app` and click `Next`.
6. Click `Show QR code`, scan it with your app and enter the `Authenticator code` to confirm.
7. Click `Done` on the confirmation screen
8. Click done and you'll appear on the AWS access portal. Click on `AdministratorAccess` underneath the account name to log into the account.

# Finished

And we're done. The lessons you will have taken away from this tutorial are:

- How AWS Organizations and IAM Identity Center work together to manage the users in your organisation and what accounts they have access to.
- That you should never log in with the root user if there are permissions you can give an IAM user.
- That permissions are best attached to groups instead of users. That way you can just look inside a group to see how many people have those particular permissions.
