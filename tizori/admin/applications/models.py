from django.db import models

# Create your models here.
class Applications(models.Model):
    id = models.TextField(primary_key=True, editable=False)
    name = models.TextField()
    credentials = models.TextField(blank=True, null=True)

    class Meta:
        db_table = 'applications'
        verbose_name = 'Application'
        verbose_name_plural = 'Applications'

    def __str__(self):
        return self.name