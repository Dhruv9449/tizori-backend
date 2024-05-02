from django.db import models

# Create your models here.
class Users(models.Model):
    username = models.TextField(primary_key=True, editable=False)
    name = models.TextField()
    email = models.TextField(blank=True, null=True)
    password = models.TextField()
    firebase_uuid = models.TextField(unique=True, blank=True, null=True)

    class Meta:
        db_table = 'users'
        verbose_name = 'User'
        verbose_name_plural = 'Users'
    
    def __str__(self):
        return self.username



