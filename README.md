HealthNear.Me
=============

**[HealthNear.Me](http://healthnear.me/)** is a tool to help regular folks find public health resources near them. This software package provides a simple, mobile-friendly website, code to index and search data, and nine City of Chicago Public Health datasets, and integration with Twilio, to allow people to interact with it via text messaging.

Background
----------

This application started as an entry in the [Making Public Health Data Work Challenge](http://www.smartchicagocollaborative.org/illinois-public-health-datapalooza-and-a-10k-challenge/). The creators, Chris Gansen and Melissa Buenger, saw that there were many datasets listing public health providers on the City of Chicago data portal, but they were all slightly different and there was no way to find all providers near a particular location.

Contributing and Reusing
------------------------

The code for HealthNear.me is released as open source software under the [MIT License](LICENSE.md). Contributors are encouraged to fork the repostory, create topic branches, and open pull requests.

The application is geared towards public health data, but with some elbow grease, it could be repurposed to index and make available most any municipal service.

Technical Overview
------------------

The application consists of a HTTP API and data loader tool, both written in Go. The web frontend is HTML and AngularJS, with Jekyll behind the scenes. ElasticSearch powers the backend datastore and search index. All of the components are hosted on Amazon Web Services (namely EC2 and S3).

Contact
-------

If you have questions about the HealthNear.me project, don't hesitate to contact us via Twitter: [@cgansen](https://twitter.com/cgansen) or [@mbuengermph](https://twitter.com/mbuengermph).