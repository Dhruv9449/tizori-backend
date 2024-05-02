from django.db import models

# Create your models here.
class Roles(models.Model):
    id = models.TextField(primary_key=True, editable=False)
    name = models.TextField(blank=True, null=True)
    permissions = models.TextField(blank=True, null=True)

    class Meta:
        db_table = 'roles'
        verbose_name = 'Role'
        verbose_name_plural = 'Roles'

    def __str__(self):
        return self.name
